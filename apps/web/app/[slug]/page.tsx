import { getWorkBySlug } from "@/services/works";
import { Title } from "@mantine/core";
import { notFound } from "next/navigation";

export default async function Works({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await params;
  const work = await getWorkBySlug(slug);

  if (!work) {
    notFound();
  }

  return (
    <div>
      <Title>{work?.title}</Title>
    </div>
  );
}
