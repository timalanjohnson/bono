import { useEffect, useState } from "react";

const connection = new WebSocket("ws://localhost:8080/ws");
connection.onopen = function (e) {
  console.log("Connection established!");
};

function App() {
  const [messages, setMessages] = useState<Array<string>>([
    "Increase front wing",
  ]);

  useEffect(() => {
    connection.onmessage = function (e) {
      console.log(e.data);
      setMessages((prev) => [...prev, e.data]);
    };
  }, []);

  return (
    <>
      <main>
        <WeatherForecast />
        <GapToLeader />
        <Session />
        <Damage />
        <PitStop />
        <Messages messages={messages} />
      </main>
      <div className="logo">
        <h1>BONO</h1>
        <span>Bespoke Optimization and Navigation Operations</span>
      </div>
    </>
  );
}

function WeatherForecast() {
  return (
    <div className="blue">
      <h2>WEATHER FORECAST:</h2>
      <ol>
        <li>
          CURRENT: <strong>DRY</strong>
        </li>
        <li>
          5 MINUTES: <strong>OVERCAST</strong>
        </li>
        <li>
          10 MINUTES: <strong>WET</strong>
        </li>
        <li>
          15 MINUTES: <strong>VERY WET</strong>
        </li>
      </ol>
    </div>
  );
}

function GapToLeader() {
  return (
    <div className="pink">
      <h2>RACE INFO</h2>
      <ul>
        <li>
          GAP TO LEADER: <strong>1:03:342</strong>
        </li>
        <li>
          GAP TO CAR AHEAD: <strong>0:01:432</strong>
        </li>
        <li>
          GAP TO CAR BEHIND: <strong>0:03:432</strong>
        </li>
      </ul>
    </div>
  );
}

function PitStop() {
  return (
    <div className="yellow">
      <h2>PIT STOP:</h2>
      <ul>
        <li>
          PIT WINDOW: <strong>LAP 10 - 12</strong>
        </li>
        <li>
          REJOIN POSITION: <strong>6TH</strong>
        </li>
      </ul>
    </div>
  );
}

function Damage() {
  return (
    <div className="red">
      <h2>DAMAGE REPORT:</h2>
      <ul>
        <li>
          FRONT WING: <strong>35%</strong>
        </li>
        <li>
          FLOOR: <strong>5%</strong>
        </li>
      </ul>
    </div>
  );
}

function Session() {
  return (
    <div className="green">
      <h2>SESSION INFO:</h2>
      <ul>
        <li>RED FLAG</li>
        <li>
          TIME REMAINING: <strong>27:34</strong>
        </li>
      </ul>
    </div>
  );
}

type MessagesProps = {
  messages: Array<string>;
};
function Messages({ messages }: MessagesProps) {
  return (
    <div className="purple">
      <h2>MESSAGES</h2>
      <ul>
        {messages.map((message) => (
          <li>{message}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
