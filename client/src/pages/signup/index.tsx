import { API_URL } from "@/constants";
import Link from "next/link";
import { useRouter } from "next/router";
import React, { useState } from "react";

const Login = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const router = useRouter();

  const signupHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();
    console.log(JSON.stringify({ email, password, username }, null, 4));
    try {
      const res = await fetch(`${API_URL}/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password, username }),
      });

      const data = await res.json();
      if (res.ok) {
        router.push("/login");
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="flex items-center justify-center min-w-full min-h-screen ">
      <div className="border border-gray-100 shadow shadow-md rounded-lg w-2/6 min-w-fit flex items-center justify-center">
        <form className="flex flex-col w-11/12 my-10">
          <div className="text-5xl font-bold text-center">
            <span className="text-blue">welcome!</span>
          </div>
          <input
            placeholder="username"
            className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
            value={username}
            autoComplete="false"
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            placeholder="email"
            className="p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
            value={email}
            autoComplete="false"
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="password"
            placeholder="password"
            className="p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
            value={password}
            autoComplete="false"
            onChange={(e) => setPassword(e.target.value)}
          />
          <button
            className="p-3 mt-6 rounded-md bg-gray-100 font-bold text-gray"
            type="submit"
            onClick={signupHandler}
          >
            Sign Up
          </button>
          <Link
            className="text-normal font-bold text-center mt-5 text-gray-600"
            href="/login"
          >
            Already have an account?
          </Link>
        </form>
      </div>
    </div>
  );
};

export default Login;
