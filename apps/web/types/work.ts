export interface Work {
  id: string;
  order: number;
  title: string;
  slug: string;
  description: string;
  url?: string;
  tags: string[];
  images: WorkImage[];
}

export interface WorkImage {
  id: string;
  filepath: string;
  workId: string;
}
