
var ws;
function Show(id){
  menus = document.getElementsByClassName("dyn_menu");
  for (var i = 0; i < menus.length; i++) {
      Hide(menus[i].id);
  }
  m=document.getElementById(id);
  m.style.display='block';
}

function Hide(id){
  m= document.getElementById(id);
  m.style.display='none';
}

function Score(type,side) {
  t=document.getElementById('event_text');
  t.value="*"+type+side+"*" + t.value;
}


function Send(){
  text=document.getElementById('event_text').value;
    ws.send(JSON.stringify({"Type":"Event","Event":{"Name":"ticker","Text":text,"Timestamp":Date.now()}}));
}

function SaveSettings(){

}

function Clear() {
  t=document.getElementById('event_text');
  t.value="";
}
function SetTeamNames(h, a) {
    e = document.getElementById('input_home_team');
    e.textContent = h;
    e = document.getElementById('input_away_team');
    e.textContent = a;
}

function SetStartTime(d) {
    e = document.getElementById('start_time');
    e.textContent = d;
}
function UpdateState(state) {
    SetTeamNames(state.HomeTeam, state.AwayTeam);
    SetStartTime(state.StartTime);

}
function initWS() {
    var socket = new WebSocket("ws://rcatickertest.hopto.org/adminws");

    socket.onopen = function () {
        document.getElementById('conn_status').innerHTML = "Connected";
    };
    socket.onmessage = function (e) {
        var json = JSON.parse(e.data);
        switch (json.Name) {
            case "State": UpdateState(json.Data);break;
            case "Event": AddEvent("Test" + new Date());
        }
    };
    socket.onclose = function () {
        document.getElementById('conn_status').innerHTML = "Diconnected";
    setTimeout(function(){start();}, 5000);
    };
    return socket;
}

function start() {
ws = initWS();
}
window.onload = start;
