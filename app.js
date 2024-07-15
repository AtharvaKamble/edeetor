async function run(textArea) {
  const code = textArea.value;

  const url = `http://localhost:8081/api/code/go`;

  try {
    fetch(url, {
      method: "POST",
      body: JSON.stringify({ code, language: "go" }),
      headers: { "Content-Type": "application/json" },
    })
      .then((resp) => resp.text())
      .then((data) => {
        // console.log(data)
        document.getElementById("cmd").innerHTML = data;
      });
  } catch (err) {
    console.error(`Error running Go code: ${err}`);
  }
}

document.onkeydown = (e) => {
  if (e.ctrlKey && e.key === `'`) {
    document.getElementById("btn-run").click();
  }
};

document
  .getElementById("editor-textarea")
  .addEventListener("keydown", function (e) {
    if (e.key == "Tab" && e.shiftKey) {
      e.preventDefault();
      var start = this.selectionStart;
      var end = this.selectionEnd;

      // set textarea value to: text before caret + tab + text after caret
      this.value =
        this.value.substring(0, start) + "\t" + this.value.substring(end);

      // put caret at right position again
      this.selectionStart = this.selectionEnd = start - 1;
    } else if (e.key == "Tab") {
      e.preventDefault();
      var start = this.selectionStart;
      var end = this.selectionEnd;

      // set textarea value to: text before caret + tab + text after caret
      this.value =
        this.value.substring(0, start) + "\t" + this.value.substring(end);

      // put caret at right position again
      this.selectionStart = this.selectionEnd = start + 1;
    }
  });
