window.addEventListener('DOMContentLoaded', () => {
    document.body.querySelector("input").focus();
    document.body.querySelector("input").addEventListener("keyup", searchBar)
    function searchBar() {
        let search = document.body.querySelector("input").value.toLowerCase();
        for (let index = 0; index < document.body.querySelectorAll("div.content").length; index++) {
            const member = document.body.querySelectorAll("div.content")[index].querySelector("p#member").innerText.toLowerCase();
            const createDate = document.body.querySelectorAll("div.content")[index].querySelector("p#createDate").innerText;
            const albumDate = document.body.querySelectorAll("div.content")[index].querySelector("p#albumDate").innerText;
            const name = document.body.querySelectorAll("div.content")[index].querySelector("h3#name").innerText.toLowerCase();

            if (member.includes(search) || albumDate.includes(search) || createDate.includes(search) || name.includes(search)) {
                document.body.querySelectorAll("div.card")[index].hidden = false;
            } else {
                document.body.querySelectorAll("div.card")[index].hidden = true;
            }
        }
    }
})
