(function() {
  let toggleForm = document.querySelector("#togglePostFormLink > a");
  toggleForm.addEventListener("click", (e) => {
    document.getElementById("togglePostFormLink").classList.add("hidden");
    document.getElementById("postForm").classList.remove("hidden");
  });
})();
