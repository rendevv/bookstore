export interface Book {
  id?: number; // dari gorm.Model
  createdAt?: string; // ISO string
  updatedAt?: string;
  deletedAt?: string | null;
  name: string;
  author: string;
  publication: string;
}
