import { Button, Flex, Title } from "@mantine/core";
import Works from "./components/Works";
import styles from "./page.module.css";
import { Work } from "@/types/work";
import { getAllWorks } from "@/services/works";
import Contact from "./components/Contact";

export default async function Home() {
  const works: Work[] = await getAllWorks();

  return (
    <Flex className="page" direction="column" align="center">
      {/* intro section */}
      <Flex
        component="section"
        id="section-intro"
        className={styles.section}
        direction="column"
        align="center"
        justify="center"
        gap="lg"
        h="100svh"
      >
        <Flex direction="column" align="center">
          <Title>
            Hi, I&apos;m{" "}
            <span style={{ color: "var(--mantine-color-primary-5)" }}>
              Ferdian
            </span>
          </Title>
          <Title>A Web Developer</Title>
        </Flex>
        <div>I build reliable and accessible systems.</div>
        <Flex gap="md">
          <Button style={{ minWidth: "10rem" }}>Contact Me</Button>
        </Flex>
      </Flex>

      <section id="section-works" className={styles.section}>
        <Works works={works} />
      </section>
      <section id="section-contact" className={styles.section}>
        <Contact />
      </section>
    </Flex>
  );
}
