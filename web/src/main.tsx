import React from "react";
import ReactDOM from "react-dom/client";
import { css, Global } from "@emotion/react";
import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";

import App from "./App";

const globalStyle = css`
  body {
    width: 100vw;
    height: 100vh;
    margin: 0;
    padding: 0;
  }

  #root {
    height: 100%;
    width: 100%;
    display: flex;
  }

  .custom-icon {
    width: 50px;
    height: 50px;
    background: transparent;
    border: none;
    position: relative;
  }
`;

const queryClient = new QueryClient({});

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Global styles={globalStyle} />
    <QueryClientProvider client={queryClient}>
      <App />
      {import.meta.env.VITE_NODE_ENV !== "prod" && <ReactQueryDevtools />}
    </QueryClientProvider>
  </React.StrictMode>
);
