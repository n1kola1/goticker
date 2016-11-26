function SetScore(h, a) {
    e = document.getElementById('home_score');
    e.textContent = h;
    e = document.getElementById('away_score');
    e.textContent = a;
}

function SetTeamNames(h, a) {
    e = document.getElementById('home_team');
    e.textContent = h;
    e = document.getElementById('away_team');
    e.textContent = a;
}

function SetStartTime(d) {
    e = document.getElementById('start_time');
    e.textContent = d;
}




function UpdateState(state) {
    SetScore(state.HomeScore, state.AwayScore);
    SetTeamNames(state.HomeTeam, state.AwayTeam);
    SetStartTime(state.StartTime);
    for (var i = 0; i < state.PreviousEvents.length; i++) {
      AddEvent(state.PreviousEvents[i]);
    }
}

function AddEvent(e) {
    t = document.getElementById('ticker');
    var li = document.createElement('li');

    if (e.Text.startsWith('*')){
      switch (e.Text[1]) {
        case 'T':  li.className = 'try';break;
        case 'C':  li.className = 'conversion';break;
        case 'P':  li.className = 'penalty'; break;
      }
      e.Text=e.Text.substring(4);

  }
  var ti=document.createElement('time');
  var time = new Date(e.Timestamp);
  ti.innerHTML=time.getHours() + ":"+('00'+time.getMinutes()).slice(-2);
//  ti.setAttribute("datetime", e.Timestamp);
  li.appendChild(ti);
    li.appendChild(document.createTextNode(e.Timestamp + " " + e.Name + " " + e.Text));
    t.insertBefore(li, t.childNodes[0]);
}
function initWS() {
    var socket = new WebSocket("ws://rcatickertest.hopto.org/tickerws");

    socket.onopen = function () {

    };
    socket.onmessage = function (e) {
        var json = JSON.parse(e.data);
        switch (json.Type) {
            case "State": UpdateState(json.State);break;
            case "Event": AddEvent(json.Event);
        }

    };
    socket.onclose = function () {
      setTimeout(function(){start();}, 5000);
    };
    return socket;
}
function start() {
    ws = initWS();
}
window.onload = start;
