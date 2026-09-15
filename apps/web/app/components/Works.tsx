import { getAllWorks } from "@/services/works";
import { Work } from "@/types/work";
import {
  Badge,
  Card,
  Divider,
  Grid,
  GridCol,
  Group,
  Text,
  Title,
} from "@mantine/core";
import styles from "@/app/page.module.css";

export default async function Works() {
  const works: Work[] = await getAllWorks();

  return (
    <div style={{ minHeight: "100vh" }}>
      <Title className={styles.sectionTitle}>Works</Title>
      <Grid align="stretch">
        {works.map((el) => (
          <GridCol key={el.id} span={6}>
            <Card withBorder h="100%" component="a" href={`/${el.slug}`}>
              <Text fw="bold">{el.title}</Text>
              <Divider my="xs" />
              <Text mb="sm">{el.description}</Text>
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
