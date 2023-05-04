import AuthContextProvider from "@/modules/auth_context";
import WebSocketProvider from "@/modules/ws_context";
import "@/styles/globals.css";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <AuthContextProvider>
      <WebSocketProvider>
        <div className="flex flex-col md:flex-row h-full min-h-screen font-sans">
          <Component {...pageProps} />
        </div>
      </WebSocketProvider>
    </AuthContextProvider>
  );
}
