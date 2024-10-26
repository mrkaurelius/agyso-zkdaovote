import { Grid, GridItem } from "@chakra-ui/react";
import BulletVoting from "./components/Voting";
import Menu from "./components/Menu";
import Wallet from "./components/Navbar";

function App() {
  return (
    <Grid
      templateAreas={`
      "nav header"
      "nav main"
      `}
      gridTemplateRows={"auto auto"}
      gridTemplateColumns={"250px auto"}
      bgGradient="linear-gradient(120deg, #a6c0fe 0%, #f68084 100%)"
      w="100vw"
      h="100vh"
      color="blackAlpha.700"
      fontWeight="bold"
    >
      <Wallet></Wallet>
      <Menu></Menu>
      <GridItem area={"main"}>
        <BulletVoting></BulletVoting>
      </GridItem>
    </Grid>
  );
}

export default App;
