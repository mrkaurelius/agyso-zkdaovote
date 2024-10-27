import {
  Box,
  Button,
  Card,
  CardBody,
  CardFooter,
  Center,
  Divider,
  FormControl,
  FormLabel,
  HStack,
  Image,
  NumberDecrementStepper,
  NumberIncrementStepper,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  useToast,
  VStack,
} from "@chakra-ui/react";

import scar from "../assets/scar.jpeg";
import apuSmart from "../assets/apu_smart.jpg";
import apuHacker from "../assets/apu_hacker.jpg";
import apuRich from "../assets/apu_rich.jpg";

import { useState } from "react";
import { getAGYSOContract } from "../lib";

function Voting() {
  const votingPower = 5;
  const [vote0, setVote0] = useState("0");
  const [vote1, setVote1] = useState("0");
  const [vote2, setVote2] = useState("0");
  const [vote3, setVote3] = useState("0");

  const [zkVoteLoading, setZkVoteLoading] = useState(false);
  const [zkVoteLoadingText, setZkVoteLoadingText] = useState("");
  const [submitLoading, setSubmitLoading] = useState(false);
  const [submitText, setSubmitText] = useState("");
  const [voteOnchainLoading, setVoteOnchainLoading] = useState(false);
  const [voteOnchainText, setVoteOnchainText] = useState("");

  const toast = useToast();

  const handleZKVoteGeneration = async () => {
    setZkVoteLoading(true);
    setZkVoteLoadingText("Generating proof for vote");

    let encryptedBallots = await getAGYSOContract().getBallotBox();
    if (encryptedBallots === "0x") encryptedBallots = "0";


    let encryptedVotesNoPrefix = encryptedBallots.slice(2, encryptedBallots.length);

    const proofBody = {
      votePower: 5,
      encryptedBallots: encryptedVotesNoPrefix,
      vote0: parseInt(vote0),
      vote1: parseInt(vote1),
      vote2: parseInt(vote2),
      vote3: parseInt(vote3),
    };

    const options = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(proofBody),
    };

    console.log(proofBody);
    let res = await fetch("http://localhost:2929/proof/vote", options);

    if (res.status < 300) {
      toast({
        position: "top",
        title: "Vote proof generated",
        status: "success",
        duration: 10000,
        isClosable: true,
      });
    } else if (res.status >= 300) {
      toast({
        position: "top",
        title: "Vote proof generation error",
        status: "error",
        isClosable: true,
      });
    }

    setZkVoteLoading(false);
  };

  const handleProofSubmittion = async () => {
    setSubmitLoading(true);
    setSubmitText("Waiting proof batch inclusion");

    const options = { method: "POST", headers: { "Content-Type": "application/json" } };
    let res = await fetch("http://localhost:2929/proof/submit", options);
    let resJson = await res.json();
    console.log(resJson);

    // const res = { status: 299 };
    // await new Promise((resolve, reject) => {
    //   setTimeout(() => {
    //     resolve(true);
    //   }, 4000);
    // });

    if (res.status < 300) {
      toast({
        position: "top",
        title: "Vote proof has been submitted and verified on the aligned layer.",
        duration: 10000,
        status: "success",
        isClosable: true,
      });
    } else if (res.status >= 300) {
      toast({
        position: "top",
        title: "Vote proof generation error",
        status: "error",
        isClosable: true,
      });
    }

    setSubmitLoading(false);
  };

  const handleOnChainVote = async () => {
    setVoteOnchainLoading(true);
    setVoteOnchainText("Submitting Vote Transaction");

    const options = { method: "GET" };
    const res = await fetch("http://localhost:2929/proof/calldata", options);
    const resBody = await res.json();
    console.log(resBody);

    const callData = resBody.calldata;
    console.log(callData);

    const contractResp = await getAGYSOContract().castZKVote(
      callData.proof_commitment_hex,
      callData.public_input_commitment_hex,
      callData.proving_system_aux_data_commitment_hex,
      callData.proof_generator_addr_hex,
      callData.batch_merkle_root_hex,
      callData.batch_inclusion_proof_hex,
      callData.index_in_batch,
      callData.pub_input_hex
    );

    await contractResp.wait(1);

    toast({
      position: "top",
      title: "Vote submitted on-chain!",
      duration: 10000,
      status: "success",
      isClosable: true,
    });

    setVoteOnchainLoading(false);
  };

  return (
    <Box margin="12" >
      <Center>
        <VStack>
          <HStack spacing="48px">
            <Card boxShadow="xl" bg="whiteAlpha.600">
              <CardBody>
                <Image src={apuHacker} w={300} h={200} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Programmer, Hacker A.</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote0}
                        onChange={(valueString) => setVote0(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
              </Center>
            </Card>

            <Card boxShadow="xl" bg="whiteAlpha.600">
              <CardBody>
                <Image src={apuSmart} w={300} h={200} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Applied Cryptographer A.</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote1}
                        onChange={(valueString) => setVote1(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
              </Center>
            </Card>

            <Card boxShadow="xl" bg="whiteAlpha.600">
              <CardBody>
                <Image src={scar} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Programmer, UI/UX 6-6 B.</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote2}
                        onChange={(valueString) => setVote2(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
              </Center>
            </Card>

            <Card boxShadow="xl" bg="whiteAlpha.600">
              <CardBody>
                <Image src={apuRich} w={300} h={200} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Dapp, DeFi, Web3 Expert N.</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote3}
                        onChange={(valueString) => setVote3(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
              </Center>
            </Card>
          </HStack>

          <Button
            mt="20"
            size="md"
            height="70px"
            width="300px"
            border="2px"
            onClick={handleZKVoteGeneration}
            isLoading={zkVoteLoading}
            loadingText={zkVoteLoadingText}
            colorScheme="blue"
          >
            Generate ZK Vote Proof
          </Button>

          <Button
            mt="5"
            size="md"
            height="70px"
            width="300px"
            border="2px"
            onClick={handleProofSubmittion}
            isLoading={submitLoading}
            loadingText={submitText}
            colorScheme="green"
          >
            Submit Vote to AlignedLayer
          </Button>

          <Button
            mt="5"
            size="md"
            height="70px"
            width="300px"
            border="2px"
            onClick={handleOnChainVote}
            isLoading={voteOnchainLoading}
            loadingText={voteOnchainText}
            colorScheme="pink"
          >
            Cast Vote On-Chain
          </Button>

        </VStack>
      </Center>
    </Box>
  );
}

export default Voting;
