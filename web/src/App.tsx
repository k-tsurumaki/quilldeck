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
      console.error('Health check failed:', error);
      if (error instanceof Error && error.message.includes('Failed to fetch')) {
        setBackendStatus('サーバーに接続中...');
        setTimeout(checkBackendHealth, 5000);
      } else {
        setBackendStatus('Backend connection failed');
      }
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
    return <AuthForm onSuccess={handleAuthSuccess} />;
  }

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Header */}
      <header className="bg-white shadow-md">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div className="flex items-center">
              <h1 className="text-3xl font-extrabold text-gray-800">QuillDeck</h1>
              <span className="ml-4 px-3 py-1 bg-green-100 text-green-800 text-sm font-medium rounded-full">
                {backendStatus.includes('ok') ? '接続中' : 'オフライン'}
              </span>
            </div>
            <div className="flex items-center space-x-4">
              <div className="text-sm text-gray-600">
                <span className="font-medium">ユーザーID:</span> {userId.substring(0, 8)}...
              </div>
              <button
                onClick={handleLogout}
                className="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg font-medium transition-colors duration-200"
              >
                ログアウト
              </button>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="space-y-8">
          {/* Upload Section */}
          <FileUpload onUploadSuccess={handleUploadSuccess} />

          {/* Documents Section */}
          {uploadedDocuments.length > 0 && (
            <div>
              <h2 className="text-2xl font-semibold text-gray-800 mb-6 flex items-center">
                <svg className="w-6 h-6 mr-3 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path>
                </svg>
                アップロード済みドキュメント ({uploadedDocuments.length})
              </h2>
              <div className="space-y-6">
                {uploadedDocuments.map((doc) => (
                  <SummaryGenerator
                    key={doc.id}
                    documentId={doc.id}
                    fileName={doc.fileName}
                  />
                ))}
              </div>
            </div>
          )}

          {/* Empty State */}
          {uploadedDocuments.length === 0 && (
            <div className="text-center py-16">
              <svg className="w-24 h-24 mx-auto text-gray-300 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="1" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <h3 className="text-xl font-medium text-gray-600 mb-2">ドキュメントをアップロードしましょう</h3>
              <p className="text-gray-500">TXTまたはMDファイルをアップロードして、AI要約を生成できます</p>
            </div>
          )}
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-white border-t border-gray-200 mt-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
          <div className="flex justify-between items-center">
            <p className="text-gray-500 text-sm">© 2025 QuillDeck. All rights reserved.</p>
            <div className="flex space-x-6 text-sm text-gray-500">
              <a href="#" className="hover:text-gray-700">ヘルプ</a>
              <a href="#" className="hover:text-gray-700">プライバシー</a>
              <a href="#" className="hover:text-gray-700">利用規約</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

export default App;