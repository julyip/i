package main

const h = `<html><head>
<meta charset="utf-8">
<style>
 body,textarea { font-family: monospace; margin:0pt; spellcheck="false";scrollbar-width:thin;scrollbar-color:#325aa8 black;}
 body { overflow: hidden; }
 textarea {background-color:black;color:white;border:none;resize: none;}
 .col { float: left; width:50%; height:100%; }
 .row:after { content: ""; display: table; clear: both }
</style>
</head><body>
<div id="dropbox">
<div class="row">
 <textarea id="term" class="col" wrap="off" spellcheck="false"></textarea>
 <textarea id="edit" class="col" wrap="off" spellcheck="false" placeholder=".e"></textarea>
 <image id="dpy" class="col"></image>
</div></div>
<script>
var term = document.getElementById("term")
var edit = document.getElementById("edit")
var dpy  = document.getElementById("dpy")
function hdr(r,h) { return r.getResponseHeader(h) }
function e(k, n, f) {
 console.log("e k",k)
 var req = new XMLHttpRequest()
 req.onreadystatechange = function() { 
  if (this.readyState == (this.DONE || 4)) { 
   // if req.getResponseHeader('Content-Type') == "image/png") {
   if (hdr(req, "n") == ".e")
    edg("/.e", hdr(req, "a"), hdr(req, "b"))
   O(req.response+" ");term.scrollTo(0, term.scrollHeight)
  } 
 }
 var a = edit.selectionStart
 var b = edit.selectionEnd
 req.open("POST", "")
 req.setRequestHeader("n", n) // file name
 req.setRequestHeader("a", a) // selection start (only .e)
 req.setRequestHeader("b", b) // selection end   (only .e)
 req.setRequestHeader("w", dpy.width) 
 req.setRequestHeader("h", dpy.height)
 req.setRequestHeader("k", k) // term value(current line)
 req.send(f)
}
function edg(u,a,b) {
 var req = new XMLHttpRequest()
 req.onreadystatechange = function() { 
  if (this.readyState == (this.DONE || 4)) { 
   if (hdr(req,"Content-Type") == "image/png") {
    dpy.src = req.response
   } else {
    edit.value = req.response
    edit.setSelectionRange(a, b)
    edit.focus()
   }
  }
 }
 req.open("GET", u)
 req.send()
}
var hold = false
term.value = " "
term.onkeydown = function (evt) {
 if (evt.which === 27) {
  evt.preventDefault()
  hold = !hold	
  term.style.border = "none"
  if (hold) 
   term.style.border = "2px solid blue"
 } else if (evt.which === 13) {
  if (hold) {
   return
  }
  evt.preventDefault()
  var a = term.selectionStart
  var b = term.selectionEnd
  var s = term.value.substring(a, b)
  if (b == a) {
   if (term.value[a] == "\n")
    a -= 1
   a = term.value.lastIndexOf("\n", a)
   if (a == -1)
    a = 0
   b = term.value.indexOf("\n", b)
   if (b == -1)
    b = term.selectionEnd
   s = term.value.substring(a, b)
  }
  if (term.selectionEnd != term.value.length)
   O(s)
  O("\n")
  s = s.trim()
  if (s === "\\c") {
   term.value = " ";return
  } else if (s === "\\e") { // toggle edit/dpy
   edit.style.display=="block"?showed(false):showed(true)
  } else {
   e(s, ".e", edit.value)
   return
  }
  P()
 }
}
function O(s) { term.value += s }
function P() { term.value += "\n "; term.scrollTo(0, term.scrollHeight) }
function show(e, b) { b?e.style.display="block":e.style.display="none" }
function showed(b) { if(b){show(edit,true);show(dpy,false)}else{show(edit,false);show(dpy,true)}}

function search3(s) { // search on button 3
 s.addEventListener("contextmenu", function(e) {
  var l = s.selectionEnd-s.selectionStart
  if (e.button==2&&l>0) {
   e.preventDefault()
   var t = s.value.substring(s.selectionStart, s.selectionEnd)
   var f = function(a) { return s.value.indexOf(t, a) }
   var n = f(s.selectionEnd)
   if (n < 0) { n = f(0) }
   s.setSelectionRange(n,n+l)
  }
 })
}
search3(term);search3(edit)


edit.addEventListener("mousedown", function(ev) { // editor: button-2(execute selection)
 console.log("which",ev.which,"button",ev.button)
 if (ev.button==1) { 
  ev.preventDefault()
  var t = edit.value.substring(edit.selectionStart, edit.selectionEnd)
  O("\n "+t+"\n");e(t,".e",edit.value)
 }
})

var dropbox = document.getElementById("dropbox")
dropbox.ondragover = function(ev) { ev.preventDefault() }
dropbox.ondrop = function(ev) {
 ev.preventDefault()
 if (ev.dataTransfer.items) {
  for (var i = 0; i< ev.dataTransfer.items.length; i++) {
   if (ev.dataTransfer.items[i].kind == 'file') {
    var file = ev.dataTransfer.items[i].getAsFile()
    addfile(file)
   }
  }
 } else
  for (var i = 0; i<ev.dataTransfer.files.length; i++)
   addfile(ev.dataTransfer.files[i])
}
function addfile(f) {
 var r = new FileReader()
 r.onload = function() {
  e("", f.name, r.result)
 }
 r.readAsArrayBuffer(f)
}
show(edit,true)
</script></body></html>`
