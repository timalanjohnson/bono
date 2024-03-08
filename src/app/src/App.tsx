import { useEffect, useState } from "react";
import "./App.css";

const connection = new WebSocket("ws://localhost:8080/ws");
connection.onopen = function (e) {
  console.log("Connection established!");
};

function App() {
  const [count, setCount] = useState(0);
  const [messages, setMessages] = useState<Array<string>>([]);

  useEffect(() => {
    connection.onmessage = function (e) {
      console.log(e.data);
      setMessages((prev) => [...prev, e.data]);
    };
  }, []);

  return (
    <>
      <h1>Bono</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
      <ul>
        {messages.map((message) => (
          <li>{message}</li>
        ))}
      </ul>
    </>
  );
}

export default App;
