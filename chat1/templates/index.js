let ws = new WebSocket("ws://localhost:8080/ws");
console.log("attemping your messages connect");
ws.onopen = () => { console.log(" connected to messages");}
ws.onclose = (event) => {console.log(" messages close: ", event);}
ws.onerror = (error) => {console.log(" messages error: ", error);}
ws.addEventListener("message", (event) => {
console.log(event);

const obj = JSON.parse(event.data)
console.log(obj.username)

if (obj.id == -1){
    document.getElementById('cont').innerHTML = '';
  } else{

const messagediv = document.createElement("div");
var user = document.createElement("h3");
var msg = document.createElement("p");

user.innerText = obj.username;
console.log(messagediv);
msg.innerText = obj.content;

messagediv.style.cssText = "padding-left: 15px;padding-top: 5px;padding-bottom: 5px; margin: 5px; background-color: rgb(192, 149, 216);";

document.getElementById('cont').appendChild(messagediv);
messagediv.appendChild(user);
messagediv.appendChild(msg);
  }
});

const wrapper = document.querySelector('.centers'),
      form    = wrapper.querySelectorAll('.center')
      submit  = form[0].querySelector('input[type=submit]');

function getdata(e){
  e.preventDefault();
  var formdata = new FormData(form[0]);
  ws.send(formdata.get('message'));
  var form1 = document.getElementById("myForm");
  form1.reset();
}

document.addEventListener('DOMContentLoaded', function(){
  submit.addEventListener('click', getdata, false )
})