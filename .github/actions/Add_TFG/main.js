const core = require("@actions/core");
// npm install node-fetch
const fetch = require("node-fetch");

const owner = "addUsername";
const repo = "victor_robles_bot";
const id = core.getInput("pull_request_id");
const token = core.getInput("repo_token");
const branch = "master";
const filename = "TFG.md";


const url_pull_request =
  "https://api.github.com/repos/" + owner + "/" + repo + "/pulls/" + id;
var merge_commit_sha;
var json;

console.log("first RESPONSE (GET PR)");
fetch(url_pull_request)
  .then((r) => {
    if (!r.ok) {
      return Promise.reject(r);
    }    
    return r.json();
  })
  .then((data) => {
    console.log("BIGG")  
    //console.log(data)
    if (!data.mergeable) {
      return Promise.reject("PR is not mergeable");
    }
    
    merge_commit_sha = data.merge_commit_sha;
    return fetch(url_pull_request + "/files");
  })
  .then((r) => {
    if (!r.ok) {
      return Promise.reject(r);
    }
    return r.json();
  })
  .then((data) => {
    json = data;
    //console.log(json);
    if (!checkPr(json)) {
      return Promise.reject(filename + " has not passed pr test");
    }
    console.log("checkPr OK");
    // Fetch another API
    const diff = json[0].patch.split("\n+");
    const columns = diff[1].split("|").map((item) => item.trim());
    if (!check_addition(columns)) {
      return Promise.reject(filename + " has not passed addition test");
    }
    console.log("check_addition OK");
    //console.log("Second RESPONSE (PING REPO)");
    //console.log(columns[2]);
    return fetch(columns[2]);
  })
  .then((r) => {
    if (!r.ok) {
      return Promise.reject(r.url + " returns: " + r.status);
    }
    console.log("check_repo OK");
    headers = getHeaders();

    return fetch(
      "https://api.github.com/repos/" + owner + "/" + repo + "/merges",
      {
        method: "POST",
        headers: headers[0],
        body: JSON.stringify(headers[1]),
      }
    );
  })
  .then((r) => {
    //console.log("Third RESPONSE (MERGE)");
    if (!r.ok) {
      console.log("not merged");
      console.log(r);
      return Promise.reject(r.url + " returns: " + r.status);
    }
    core.setOutput("output_mssg", "cool");
    return;
  })
  .catch(function (error) {
    core.setOutput("output_mssg", error);
    core.setFailed(error);
  });

function checkPr(json) {
  if (json.length > 1) {
    console.log("more than 1 file was modified");
    return false;
  }
  if (json[0].filename.localeCompare(filename) != 0) {
    console.log(filename + " not present");
    return false;
  }
  if (json[0].deletions > 0) {
    console.log(filename + " has deletions");
    return false;
  }
  if (json[0].additions != 1) {
    console.log(filename + " has too many additions");
    return false;
  }
  return true;
}
function check_addition(columns) {
  //| Title | Repo | Description |
  // wrong nums of columns
  if (columns.length != 5) {
    console.log("worng length columns");
    return false;
  }
  // bad description length
  if (columns[3].length > 100 || columns[3].length < 7) {
    console.log("wrong length");
    return false;
  }
  var regex = new RegExp(/https:\/\/(www\.)?github\.com\/.*$/gi);
  
  if (!columns[2].match(regex)) {
    console.log(columns[2] + " is not a github repo");
    return false;
  } 
  return true;
}
function getHeaders() {
  const headers = {
    authorization: "Bearer " + token,
    "content-type": "application/json",
    Accept: " application/vnd.github.v3+json",
  };
  const data = {
    base: branch,
    head: merge_commit_sha,
  };
  return [headers, data];
}
