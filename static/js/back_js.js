document.addEventListener("DOMContentLoaded", (e) => {
  document.querySelectorAll(".toast").forEach((elm) => new bootstrap.Toast(elm).show());
});
