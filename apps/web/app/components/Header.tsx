"use client";

import {
  ActionIcon,
  Anchor,
  Flex,
  Switch,
  Title,
  useMantineColorScheme,
} from "@mantine/core";
import { MoonIcon, SunIcon } from "@phosphor-icons/react";

export default function Header() {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();

  const handleScroll = (id: string) => {
    document.getElementById(id)?.scrollIntoView({
      behavior: "smooth",
    });
  };

  return (
    <Flex
      pos="fixed"
      top="16px"
      left={0}
      right={0}
      align="center"
      justify="space-between"
      gap="lg"
      h="46px"
      p="1.5rem 2rem"
      style={{
        maxWidth: "500px",
        margin: "0 auto",
        borderRadius: "30px",
        border: "1px solid var(--mantine-color-primary-4)",
        backgroundColor: "rgba(242, 242, 242, 0.8)",
        backdropFilter: "blur(4px)",
        WebkitBackdropFilter: "blur(4px)",
        zIndex: 99,
      }}
    >
      <Anchor onClick={() => handleScroll("section-intro")} w="76px">
        <Title order={3}>Ianz</Title>
      </Anchor>
      <Flex gap="md">
        <Anchor onClick={() => handleScroll("section-works")}>Works</Anchor>
        <Anchor onClick={() => handleScroll("section-skills")}>Skills</Anchor>
        <Anchor onClick={() => handleScroll("section-contact")}>Contact</Anchor>
      </Flex>
      <Flex justify="end" w="76px">
        <ActionIcon variant="light" onClick={toggleColorScheme}>
          {colorScheme === "dark" ? <SunIcon /> : <MoonIcon />}
        </ActionIcon>
      </Flex>
    </Flex>
  );
}
