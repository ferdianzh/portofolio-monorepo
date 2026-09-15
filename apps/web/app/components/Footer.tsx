import { Center, Text } from "@mantine/core";

export default function Footer() {
  const currentYear = new Date().getFullYear();

  return (
    <Center py="xs">
      <Text>@{currentYear} by Ferdianzh</Text>
    </Center>
  );
}
