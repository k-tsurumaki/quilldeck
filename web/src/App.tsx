import React, { useState, useEffect } from 'react';
import { AuthForm } from './components/AuthForm';
import { FileUpload } from './components/FileUpload';
import { SummaryGenerator } from './components/SummaryGenerator';
import { api } from './api/client';

interface UploadedDocument {
  id: string;
  fileName: string;
}

function App() {
  const [userId, setUserId] = useState<string>('');
  const [backendStatus, setBackendStatus] = useState<string>('');
  const [uploadedDocuments, setUploadedDocuments] = useState<UploadedDocument[]>([]);

  useEffect(() => {
    checkBackendHealth();
  }, []);

  const checkBackendHealth = async () => {
    try {
      const result = await api.health();
      setBackendStatus(`Backend Status: ${result.status}`);
    } catch (error) {
      setBackendStatus('Backend connection failed');
    }
  };

  const handleAuthSuccess = (newUserId: string) => {
    setUserId(newUserId);
  };

  const handleUploadSuccess = (documentId: string, fileName: string) => {
    setUploadedDocuments(prev => [...prev, { id: documentId, fileName }]);
  };

  const handleLogout = () => {
    setUserId('');
    setUploadedDocuments([]);
  };

  if (!userId) {
    return (
      <div style={{ minHeight: '100vh', backgroundColor: '#f8f9fa' }}>
        <div style={{ padding: '20px', textAlign: 'center' }}>
          <h1 style={{ color: '#333', marginBottom: '10px' }}>QuillDeck</h1>
          <p style={{ color: '#666', marginBottom: '20px' }}>
            ドキュメント要約アプリケーション
          </p>
          <p style={{ fontSize: '14px', color: '#28a745' }}>
            {backendStatus}
          </p>
        </div>
        <AuthForm onSuccess={handleAuthSuccess} />
      </div>
    );
  }

  return (
    <div style={{ minHeight: '100vh', backgroundColor: '#f8f9fa' }}>
      <header style={{ 
        backgroundColor: 'white', 
        padding: '15px 20px', 
        borderBottom: '1px solid #ddd',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center'
      }}>
        <h1 style={{ margin: 0, color: '#333' }}>QuillDeck</h1>
        <div>
          <span style={{ marginRight: '15px', color: '#666' }}>
            User ID: {userId.substring(0, 8)}...
          </span>
          <button
            onClick={handleLogout}
            style={{
              padding: '8px 16px',
              backgroundColor: '#dc3545',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            ログアウト
          </button>
        </div>
      </header>

      <main style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
        <FileUpload onUploadSuccess={handleUploadSuccess} />

        {uploadedDocuments.length > 0 && (
          <div>
            <h2>アップロード済みドキュメント</h2>
            {uploadedDocuments.map((doc) => (
              <SummaryGenerator
                key={doc.id}
                documentId={doc.id}
                fileName={doc.fileName}
              />
            ))}
          </div>
        )}

        {uploadedDocuments.length === 0 && (
          <div style={{ 
            textAlign: 'center', 
            padding: '40px', 
            color: '#666' 
          }}>
            <p>ファイルをアップロードして要約を生成しましょう</p>
          </div>
        )}
      </main>
    </div>
  );
}

export default App;