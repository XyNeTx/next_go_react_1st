'use client'
import Image from "next/image";
import {JSX, useEffect, useState } from "react";
import Click from "../components/click";
import axios from "axios";

type GetResponse = {
  Avatar: string | null,
  CreateAt: string | null,
  CreateBy: string | null,
  Email: string | null,
  LastLogin: string | null,
  NameJP: string | null,
  NameTH: string | null,
  ResetToken: string | null,
  Status: string | null,
  SupplierCode: string | null,
  SurnameJP: string | null,
  SurnameTH: string | null,
  TitleJP_ID: string | null,
  TitleTH_ID: string | null,
  Token: string | null,
  UIExpandIcon: string | null,
  UIHeader: string | null,
  UIHeaderBrand: string | null,
  UIIconColor: string | null,
  UILanguage: string | null,
  UILinkColor: string | null,
  UIMenuColor: string | null,
  UIMenuIcon: string | null,
  UISideBar: string | null,
  UITheme: string | null,
  UpdateAt: string | null,
  UpdateBy: string | null,
  code: string | null,
  name: string | null,
  password: string | null,
  surname: string | null,
  title_id: string | null
};


export function Testget() : JSX.Element {
  const [get, setGet] = useState<GetResponse | null>(null);
  useEffect(() =>{
    fetchData();
  },[]);

  function fetchData() {
    axios
      .get("http://localhost:8080/")
      .then((response) => {
        console.log(response);
        setGet(response.data.message);
        console.log(get);
      })
      .catch((error) => {
        console.log(error);
        return <p>Error fetching data</p>;
      });
  }

  return (
          <>
            {get !== null ? (
              <table style={{ minWidth: "600px" }} className="table-auto border border-slate-400 border-collapse">
                <thead className="text-center">
                <tr>
                <th>Code</th>
                <th>Name</th>
                <th>Surname</th>
                <th>Email</th>
              </tr>
            </thead>
            <tbody className="text-center">
              <tr>
                <td>{get?.code}</td>
                <td>{get?.name}</td>
                <td>{get?.surname}</td>
                <td>{get?.Email}</td>
              </tr>
            </tbody>
          </table>
        ) : (
          <p>No data available</p>
        )}
      </>
    );
}


export default function Home() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <Testget/>
      <main className="flex flex-col gap-[32px] row-start-2 items-center sm:items-start">
        <Image
          className="dark:invert"
          src="/next.svg"
          alt="Next.js logo"
          width={180}
          height={38}
          priority
        />
        {/* <Title persons={PersonData()} /> */}
        <Click />
        <ol className="list-inside list-decimal text-sm/6 text-center sm:text-left font-[family-name:var(--font-geist-mono)]">
          <li className="mb-2 tracking-[-.01em]">
            Get started by editing{" "}
            <code className="bg-black/[.05] dark:bg-white/[.06] px-1 py-0.5 rounded font-[family-name:var(--font-geist-mono)] font-semibold">
              src/app/page.tsx
            </code>
            .
          </li>
          <li className="tracking-[-.01em]">
            Save and see your changes instantly.
          </li>
        </ol>

        <div className="flex gap-4 items-center flex-col sm:flex-row">
          <a
            className="rounded-full border border-solid border-transparent transition-colors flex items-center justify-center bg-foreground text-background gap-2 hover:bg-[#383838] dark:hover:bg-[#ccc] font-medium text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:w-auto"
            href="https://vercel.com/new?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
            target="_blank"
            rel="noopener noreferrer"
          >
            <Image
              className="dark:invert"
              src="/vercel.svg"
              alt="Vercel logomark"
              width={20}
              height={20}
            />
            Deploy now
          </a>
          <a
            className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent font-medium text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 w-full sm:w-auto md:w-[158px]"
            href="https://nextjs.org/docs?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
            target="_blank"
            rel="noopener noreferrer"
          >
            Read our docs
          </a>
        </div>
      </main>
      <footer className="row-start-3 flex gap-[24px] flex-wrap items-center justify-center">
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://nextjs.org/learn?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image
            aria-hidden
            src="/file.svg"
            alt="File icon"
            width={16}
            height={16}
          />
          Learn
        </a>
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://vercel.com/templates?framework=next.js&utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image
            aria-hidden
            src="/window.svg"
            alt="Window icon"
            width={16}
            height={16}
          />
          Examples
        </a>
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://nextjs.org?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image
            aria-hidden
            src="/globe.svg"
            alt="Globe icon"
            width={16}
            height={16}
          />
          Go to nextjs.org →
        </a>
      </footer>
    </div>
  );
}
