import React, { useState } from 'react';
import { api } from '../api/client';

interface FileUploadProps {
  onUploadSuccess: (documentId: string, fileName: string) => void;
}

export const FileUpload: React.FC<FileUploadProps> = ({ onUploadSuccess }) => {
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [error, setError] = useState('');

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      const fileExtension = selectedFile.name.toLowerCase();
      if (fileExtension.endsWith('.txt') || fileExtension.endsWith('.md')) {
        setFile(selectedFile);
        setError('');
      } else {
        setError('TXTまたはMDファイルのみアップロード可能です');
        setFile(null);
      }
    }
  };

  const handleUpload = async () => {
    if (!file) return;

    setUploading(true);
    setError('');

    try {
      const result = await api.uploadDocument(file);
      
      if (result.error) {
        setError(result.error);
      } else {
        onUploadSuccess(result.document_id, file.name);
        setFile(null);
        // ファイル入力をリセット
        const fileInput = document.getElementById('file-input') as HTMLInputElement;
        if (fileInput) fileInput.value = '';
      }
    } catch (err) {
      setError('アップロードに失敗しました');
    } finally {
      setUploading(false);
    }
  };

  return (
    <div style={{ 
      border: '2px dashed #ccc', 
      borderRadius: '8px', 
      padding: '20px', 
      textAlign: 'center',
      marginBottom: '20px'
    }}>
      <h3>ファイルアップロード</h3>
      
      <input
        id="file-input"
        type="file"
        accept=".txt,.md"
        onChange={handleFileChange}
        style={{ marginBottom: '15px' }}
      />

      {file && (
        <div style={{ marginBottom: '15px' }}>
          <p>選択されたファイル: <strong>{file.name}</strong></p>
          <p>サイズ: {(file.size / 1024).toFixed(2)} KB</p>
        </div>
      )}

      {error && (
        <div style={{ color: 'red', marginBottom: '15px' }}>
          {error}
        </div>
      )}

      <button
        onClick={handleUpload}
        disabled={!file || uploading}
        style={{
          padding: '10px 20px',
          backgroundColor: file && !uploading ? '#28a745' : '#ccc',
          color: 'white',
          border: 'none',
          borderRadius: '4px',
          cursor: file && !uploading ? 'pointer' : 'not-allowed'
        }}
      >
        {uploading ? 'アップロード中...' : 'アップロード'}
      </button>

      <p style={{ fontSize: '14px', color: '#666', marginTop: '10px' }}>
        対応形式: .txt, .md (最大10MB)
      </p>
    </div>
  );
};