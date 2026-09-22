import { Anchor, Flex, Title } from "@mantine/core";

export default function Header() {
  return (
    <Flex
      pos="fixed"
      top="2rem"
      left={0}
      right={0}
      align="center"
      justify="space-between"
      gap="lg"
      h="46px"
      p="1.5rem 3rem"
      style={{
        maxWidth: "50vw",
        margin: "0 auto",
        borderRadius: "30px",
        border: "1px solid var(--mantine-color-primary-4)",
        backgroundColor: "rgba(242, 242, 242, 0.8)",
        backdropFilter: "blur(4px)",
        WebkitBackdropFilter: "blur(4px)",
        zIndex: 99,
      }}
    >
      <Anchor href="/">
        <Title order={3}>Ferdian</Title>
      </Anchor>
      <Flex gap="xl">
        <Anchor href="/">Works</Anchor>
        <Anchor href="/">Skills</Anchor>
        <Anchor href="/">Contact</Anchor>
      </Flex>
    </Flex>
  );
}
