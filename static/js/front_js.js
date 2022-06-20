(function() {
  let toggleForm = document.querySelector("#togglePostFormLink > a");
  toggleForm.addEventListener("click", (e) => {
    document.getElementById("togglePostFormLink").classList.add("hidden");
    document.getElementById("postForm").classList.remove("hidden");
  });

  let thread = document.querySelector(".thread");
  thread.addEventListener("click", (e) => {
    let post = e.target;
    if (post.classList.contains("quotePost")) {
      let pid = post.dataset.id;
      document.getElementById("qrTid").innerHTML = pid;
      document.getElementById("quickReply").classList.remove("hidden");
      document.querySelector("#quickReply .content").innerHTML = ">>" + pid + "\n";
    }
  });
})();
