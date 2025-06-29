import React, { useState } from 'react';
import { api } from '../api/client';

interface FileUploadProps {
  onUploadSuccess: (documentId: string, fileName: string) => void;
}

export const FileUpload: React.FC<FileUploadProps> = ({ onUploadSuccess }) => {
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [uploadProgress, setUploadProgress] = useState(0);
  const [error, setError] = useState('');

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      const fileExtension = selectedFile.name.toLowerCase();
      if (fileExtension.endsWith('.txt') || fileExtension.endsWith('.md')) {
        setFile(selectedFile);
        setError('');
        setUploadProgress(0);
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
    setUploadProgress(0);

    // プログレスバーのシミュレーション
    const progressInterval = setInterval(() => {
      setUploadProgress(prev => {
        if (prev >= 90) {
          clearInterval(progressInterval);
          return 90;
        }
        return prev + 10;
      });
    }, 100);

    try {
      const result = await api.uploadDocument(file);
      
      clearInterval(progressInterval);
      setUploadProgress(100);
      
      if (result.error) {
        setError(result.error);
      } else {
        onUploadSuccess(result.document_id, file.name);
        setFile(null);
        const fileInput = document.getElementById('file-input') as HTMLInputElement;
        if (fileInput) fileInput.value = '';
        setTimeout(() => setUploadProgress(0), 2000);
      }
    } catch (err) {
      clearInterval(progressInterval);
      console.error('Upload error:', err);
      if (err instanceof Error) {
        if (err.message.includes('Failed to fetch')) {
          setError('サーバーに接続できません。しばらくお待ちください。');
        } else if (err.message.includes('HTTP error')) {
          setError('アップロードに失敗しました。サーバーエラーが発生しました。');
        } else {
          setError('アップロードに失敗しました。ネットワークエラーが発生しました。');
        }
      } else {
        setError('アップロードに失敗しました。');
      }
      setUploadProgress(0);
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="bg-white p-8 rounded-lg shadow-md">
      <h2 className="text-2xl font-semibold mb-6 text-gray-700">資料アップロード</h2>
      
      <div className="border-2 border-dashed border-gray-300 rounded-lg p-12 text-center mb-8 bg-gray-50 hover:border-indigo-400 transition-all duration-200">
        <input
          id="file-input"
          type="file"
          accept=".txt,.md"
          onChange={handleFileChange}
          className="hidden"
        />
        <label htmlFor="file-input" className="cursor-pointer">
          <svg className="w-16 h-16 mx-auto text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="1" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a3 3 0 013 3v10a2 2 0 01-2 2H7a2 2 0 01-2-2V16m2-2l4-4m0 0l4 4m-4-4v10"></path>
          </svg>
          <p className="text-gray-600 text-lg">ファイルをドラッグ＆ドロップするか、<span className="text-indigo-600 font-medium">クリックして選択</span></p>
          <p className="text-gray-500 text-sm mt-2">対応形式: .txt, .md (最大10MB)</p>
        </label>
      </div>

      {file && (
        <div className="bg-white border border-gray-200 rounded-lg p-6 shadow-sm mb-6">
          <p className="text-gray-700 font-medium mb-3">選択されたファイル:</p>
          <div className="flex items-center mb-2">
            <span className="text-indigo-600 font-semibold text-lg mr-4">{file.name}</span>
            <span className="text-gray-500 text-sm">({(file.size / 1024).toFixed(2)} KB)</span>
          </div>
          
          {uploadProgress > 0 && (
            <>
              <div className="flex items-center mb-2">
                <span className="text-gray-500 text-sm">{uploadProgress}%</span>
              </div>
              <div className="w-full bg-gray-200 rounded-full h-2.5">
                <div
                  className="bg-indigo-500 h-2.5 rounded-full transition-all duration-300"
                  style={{ width: `${uploadProgress}%` }}
                ></div>
              </div>
              {uploadProgress === 100 && (
                <p className="text-green-600 text-sm mt-3 flex items-center">
                  <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                  アップロード完了！要約生成に進めます。
                </p>
              )}
            </>
          )}
        </div>
      )}

      {error && (
        <div className="mb-6 p-3 bg-red-100 border border-red-400 text-red-700 rounded-lg">
          {error}
        </div>
      )}

      <button
        onClick={handleUpload}
        disabled={!file || uploading}
        className={`w-full py-3 px-6 rounded-lg font-semibold transition-colors duration-200 shadow-md ${
          file && !uploading
            ? 'bg-indigo-600 hover:bg-indigo-700 text-white'
            : 'bg-gray-300 cursor-not-allowed text-gray-500'
        }`}
      >
        {uploading ? 'アップロード中...' : 'ファイルをアップロード'}
      </button>
    </div>
  );
};