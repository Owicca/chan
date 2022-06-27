'use strict';
document.addEventListener("DOMContentLoaded", (e) => {
  document.querySelectorAll(".toast").forEach((elm) => new bootstrap.Toast(elm).show());

  let role = document.getElementById("role");
  if (role) {
    role.addEventListener("change", function(e) {
      let optionList = e.target.options;
      let option = optionList[optionList.selectedIndex];

      let board = document.getElementById("boardCnt");
      switch(option.text) {
        case "board_admin":
        case "op":
            board.classList.remove("d-none");
          break;
        default:
            board.classList.add("d-none");
          break;
      }
    });
  }
});
