"use client";

import {
  Button,
  Card,
  Group,
  Text,
  Textarea,
  TextInput,
  Title,
} from "@mantine/core";
import styles from "@/app/page.module.css";
import { useForm } from "@mantine/form";

export default function Contact() {
  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      email: "",
      message: "",
    },

    validate: {
      name: (value) => (value.length > 0 ? null : "Name is required"),
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      message: (value) => (value.length > 0 ? null : "Message is required"),
    },
  });

  return (
    <div style={{ minHeight: "100svh" }}>
      <Title className={styles.sectionTitle}>Contact</Title>

      <Card withBorder w={{ base: "100%", sm: "620px" }} m="auto">
        <form onSubmit={form.onSubmit((values) => console.log(values))}>
          <TextInput
            label="Name"
            placeholder="Your Name"
            key={form.key("name")}
            {...form.getInputProps("name")}
          />

          <TextInput
            label="Email"
            placeholder="your@email.com"
            key={form.key("email")}
            {...form.getInputProps("email")}
          />

          <Textarea
            label="Message"
            placeholder="Write your message"
            key={form.key("message")}
            autosize
            minRows={4}
            {...form.getInputProps("message")}
          />

          <Group justify="flex-end" mt="md">
            <Button type="submit">Submit</Button>
          </Group>
        </form>
      </Card>

      <Text component="footer" ta="center" mt="48px">
        @{new Date().getFullYear()} Muhammad Ferdiansyah
      </Text>
    </div>
  );
}
