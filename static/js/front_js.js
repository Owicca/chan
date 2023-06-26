'use strict';
(function() {
  let threadListPostFormLink = document.querySelector("#togglePostFormLink > a");
  if(threadListPostFormLink) {
    threadListPostFormLink.addEventListener("click", openPostForm);
  }
  let postFormLink = document.querySelector(".open-qr-link");
  if(postFormLink) {
    postFormLink.addEventListener("click", openPostForm);
  }

  let thread = document.querySelector(".thread");

  /**
   * Show quick form
   */
  thread && thread.addEventListener("click", (e) => {
    console.log("here");
    let post = e.target;
    if (!post.classList.contains("quotePost")) {
      return;
    }

    let qr = document.getElementById("quickReply");
    if (!qr.classList.contains("hidden")) {
      return;
    }

    let pid = post.dataset.id;
    let form = document.getElementById("quickReplyForm");
    let boardCode = form.dataset.boardCode;
    form.dataset.curpid = pid;
    document.getElementById("qrTid").innerHTML = pid;

    qr.classList.remove("hidden");
  });

  /**
   * Add link
   */
  thread && thread.addEventListener("click", (e) => {
    console.log("here2");
    let post = e.target;
    if (!post.classList.contains("quotePost")) {
      return;
    }

    let pid = post.dataset.id;
    let form = document.getElementById("quickReplyForm");
    let boardCode = form.dataset.boardCode;
    addLink(boardCode, pid, form.dataset.curpid);
  });

  let close = document.getElementById("qrClose");
  close && close.addEventListener("click", function(e) {
    clearFormLinks();
    document.getElementById("quickReply").classList.add("hidden");
  });
})();

function openPostForm() {
  document.getElementById("togglePostFormLink").classList.add("hidden");
  document.getElementById("postForm").classList.remove("hidden");
};

function createLinkSpan(boardCode, pid) {
  let el = document.createElement("span");
  el.href = "/boards/${boardCode}/threads/${pid}/";
  el.classList.add("quoteLink");
  el.innerHTML = ">>" + pid;

  return el;
}

function createLinkInput(boardCode, pid) {
  let newInp = document.createElement("input");
  newInp.setAttribute("name", "links[]");
  newInp.setAttribute("type", "hidden");
  newInp.setAttribute("value", pid);

  return newInp;
}

function clearFormLinks() {
  document.querySelectorAll('.quoteLink').forEach(el => {
    el.remove();
  });
}

function addLink(boardCode, pid, curPid) {
  let data = new Set();

  document.querySelectorAll('input[name="links[]"]').forEach(el => {
    data.add(el.value);
  });

  if (data.has(pid)) {
    return
  }

  clearFormLinks();
  if(curPid != null && curPid != pid) {
    data.add(pid);
  }

  data.forEach(v => {
    let inp = createLinkInput(boardCode, v);
    let link = createLinkSpan(boardCode, v);
    link.appendChild(inp);

    let linkCnt = document.getElementById("linkCnt")
    linkCnt.appendChild(link);
  });
}
