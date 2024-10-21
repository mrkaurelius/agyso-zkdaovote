import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const AGYSODaoVoteValidator = buildModule(

  "AGYSODaoVoteValidatorModule",
  (m) => {
    const ALIGNED_SERVICE_MANAGER_ADDRESS =
      "0x58f280bebe9b34c9939c3c39e0890c81f163b623";
    const PAYMENT_SERVICE_ADDRESS =
      "0x815aeca64a974297942d2bbf034abee22a38a003";

    const validator = m.contract(
      "AGYSODaoVoteValidator",
      [ALIGNED_SERVICE_MANAGER_ADDRESS, PAYMENT_SERVICE_ADDRESS],
      {}
    );

    return { validator };
  }
);

export default AGYSODaoVoteValidator;
