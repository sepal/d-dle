import type { NextPage } from "next";
import Head from "next/head";
import styled from "styled-components";
import Image from "next/image";
import Game from "../components/Game";
import { GameData } from "../models/game_manager";

const Title = styled.h1`
  max-width: 256px;
  font-size: 2em;
  margin: 0 0 0 0;
  color: #333;
  text-align: center;
`;

const Header = styled.header`
  display: flex;
  justify-content: center;
  align-items: center;
  max-width: 512px;
  margin: 0.5em auto;
`;

type HomeProps = {
  game: GameData;
};

const Home: NextPage<HomeProps> = ({ game }: HomeProps) => {
  return (
    <div>
      <Head>
        <title>Dodle</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header>
        <Image
          src="/android-chrome-192x192.png"
          alt="dodle"
          width={24}
          height={24}
          layout="fixed"
        />
        <Title>Dodle</Title>
      </Header>
      <main>
        <Game game={game} />
      </main>

      <footer></footer>
    </div>
  );
};

export async function getServerSideProps() {
  const url = process.env.VERCEL_URL
    ? `https://${process.env.VERCEL_URL}`
    : "http://localhost:3000";

  const res = await fetch(`${url}/api/game/`);
  const game = await res.json();

  return { props: { game: game } };
}

export default Home;
