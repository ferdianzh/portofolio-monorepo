"use client";

import { Work } from "@/types/work";
import {
  Badge,
  Card,
  Divider,
  Grid,
  GridCol,
  Group,
  Image,
  Text,
  Title,
} from "@mantine/core";
import styles from "@/app/page.module.css";

export default function Works({ works }: { works: Work[] }) {
  return (
    <div>
      <Title className={styles.sectionTitle}>Works</Title>
      <Grid align="stretch" gap="xl">
        {works.map((el) => (
          <GridCol key={el.id} span={6}>
            <Card withBorder h="100%" component="a" href={`/${el.slug}`}>
              <Card.Section>
                <Image
                  height={200}
                  alt={el.title}
                  fallbackSrc="https://placehold.co/400x200?text=Placeholder"
                />
              </Card.Section>

              <Text fw="bold" my="sm">
                {el.title}
              </Text>
              <Text mb="lg">{el.description}</Text>
              <Group mt="auto" gap="sm">
                {el.tags.map((tag, i) => (
                  <Badge variant="light" radius="sm" key={i}>
                    {tag}
                  </Badge>
                ))}
              </Group>
            </Card>
          </GridCol>
        ))}
      </Grid>
    </div>
  );
}
