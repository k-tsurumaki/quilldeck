const API_BASE = 'http://localhost:8080/api';

export interface User {
  id: string;
  email: string;
  name: string;
}

export interface Document {
  id: string;
  title: string;
  type: string;
  size: number;
  uploaded_at: string;
}

export interface Summary {
  id: string;
  content: string;
  keywords: string[];
}

export const api = {
  // 認証
  register: async (email: string, password: string, name: string) => {
    const response = await fetch(`${API_BASE}/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password, name }),
    });
    return response.json();
  },

  login: async (email: string, password: string) => {
    const response = await fetch(`${API_BASE}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    return response.json();
  },

  // ドキュメント
  uploadDocument: async (file: File) => {
    const formData = new FormData();
    formData.append('file', file);
    
    const response = await fetch(`${API_BASE}/documents/upload`, {
      method: 'POST',
      body: formData,
    });
    return response.json();
  },

  generateSummary: async (documentId: string, length: 'short' | 'medium' | 'long') => {
    const response = await fetch(`${API_BASE}/documents/summary`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ document_id: documentId, length }),
    });
    return response.json();
  },

  // ヘルスチェック
  health: async () => {
    const response = await fetch('http://localhost:8080/health');
    return response.json();
  },
};