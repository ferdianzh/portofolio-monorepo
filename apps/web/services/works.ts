import { Work } from "@/types/work";

export async function getAllWorks(): Promise<Work[]> {
  return dummyWorks;
}

export async function getWorkBySlug(slug: string): Promise<Work | null> {
  return dummyWorks.find((work) => work.slug === slug) || null;
}

const dummyWorks: Work[] = [
  {
    id: "588a13cb-0282-451a-9c78-77c28afafb92",
    order: 1,
    title: "Lajureaksi",
    slug: "lajureaksi",
    description:
      "A commercial e-learning platform for environmental engineering.",
    url: "https://example.com",
    tags: ["SQL", "TypeScript", "NestJS", "React"],
    images: [
      {
        id: "4d8dc0c2-d236-4d56-b427-3fcb1c076b9c",
        filepath: "/images/work1-image1.jpg",
        workId: "588a13cb-0282-451a-9c78-77c28afafb92",
      },
      {
        id: "98e0d13d-e1d9-4c66-aaeb-c713d266772a",
        filepath: "/images/work1-image2.jpg",
        workId: "588a13cb-0282-451a-9c78-77c28afafb92",
      },
    ],
  },
  {
    id: "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed",
    order: 2,
    title: "Internal LMS",
    slug: "internal-lms",
    description: "Internal learning management system for company.",
    url: "https://example.com",
    tags: ["SQL", "TypeScript", "NestJS", "React"],
    images: [
      {
        id: "f7b60685-c1aa-473e-ae8d-d95970248a80",
        filepath: "/images/work2-image1.jpg",
        workId: "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed",
      },
      {
        id: "3c306c83-d422-44a9-9a66-c2fbc76420a6",
        filepath: "/images/work2-image2.jpg",
        workId: "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed",
      },
    ],
  },
];
