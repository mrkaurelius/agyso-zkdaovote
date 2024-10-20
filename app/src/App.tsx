import { useEffect } from "react";
import { Box } from "@chakra-ui/react";
import ConnectWallet from "./components/ConnectWallet";
import BulletVoting from "./components/BulletVoting";
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
    <Box>
      <ConnectWallet></ConnectWallet>
      <Box>
        <BulletVoting></BulletVoting>
      </Box>
    </Box>
  );
}

export default App;
