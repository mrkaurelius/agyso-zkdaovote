import { useEffect } from "react";

const go = new Go();

function App() {
  useEffect(() => {
    WebAssembly.instantiateStreaming(fetch("zk.wasm"), go.importObject).then(
      (result) => {
        go.run(result.instance);
      }
    );
  });

  return (
    <>
      <div>
        <h1>Merhaba Yalan Dunya!</h1>
      </div>
    </>
  );
}

export default App;
