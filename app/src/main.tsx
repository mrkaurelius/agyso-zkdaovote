import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ChakraProvider } from "@chakra-ui/react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import VotePage from "./pages/VotePage.tsx";
import ResultsPage from "./pages/ResultsPage.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <VotePage></VotePage>,
  },
  {
    path: "/results",
    element: <ResultsPage></ResultsPage>,
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ChakraProvider>
      <RouterProvider router={router} />
    </ChakraProvider>
  </StrictMode>
);
