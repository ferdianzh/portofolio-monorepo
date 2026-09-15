import { Button, Flex, Title } from "@mantine/core";
import Works from "./components/Works";
import Skills from "./components/Skills";
import styles from "./page.module.css";

export default function Home() {
  return (
    <Flex className="page" direction="column" gap="xl" align="center">
      {/* intro section */}
      <Flex
        component="section"
        className={styles.section}
        direction="column"
        align="center"
        gap="lg"
        pt="8rem"
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

      <section className={styles.section}>
        <Works />
      </section>
      <section className={styles.section}>
        <Skills />
      </section>
    </Flex>
  );
}
