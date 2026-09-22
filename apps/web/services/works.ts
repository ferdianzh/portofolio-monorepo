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
    tags: ["MySQL", "TypeScript", "NestJS", "React"],
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
    url: undefined,
    tags: ["MySQL", "TypeScript", "NestJS", "React"],
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
  {
    id: "2a47207a-37d4-43f2-a197-de1b1bc0c6e0",
    order: 3,
    title: "Benchmark Analytics App",
    slug: "benchmark-analytics-app",
    description: "A web application to display and analyze data.",
    url: undefined,
    tags: ["MySQL", "TypeScript", "NestJS", "React"],
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
  {
    id: "073672c2-ae4a-4c9f-a554-d24049771180",
    order: 4,
    title: "Service Order Management",
    slug: "service-order-management",
    description:
      "A web application for customers to create and track service orders.",
    url: undefined,
    tags: ["MySQL", "PHP", "Laravel", "Vue"],
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
  {
    id: "d1e3ea9a-4877-48b0-941a-e351f6a5e813",
    order: 5,
    title: "Warehouse Loading System",
    slug: "warehouse-loading-system",
    description: "A web application to manage warehouse loading process.",
    url: undefined,
    tags: ["PostgreSQL", "TypeScript", "Express", "React"],
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
  {
    id: "694387a9-662d-4a61-87da-7f51a279a6ce",
    order: 6,
    title: "Online Course Admin Dashboard",
    slug: "online-course-admin-dashboard",
    description: "A web application for admin to manage online courses.",
    url: undefined,
    tags: ["MySQL", "PHP", "CodeIgniter"],
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
