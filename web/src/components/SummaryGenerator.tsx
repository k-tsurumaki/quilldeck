import React, { useState } from 'react';
import { api } from '../api/client';

interface SummaryGeneratorProps {
  documentId: string;
  fileName: string;
}

export const SummaryGenerator: React.FC<SummaryGeneratorProps> = ({ 
  documentId, 
  fileName 
}) => {
  const [length, setLength] = useState<'short' | 'medium' | 'long'>('short');
  const [summary, setSummary] = useState<string>('');
  const [keywords, setKeywords] = useState<string[]>([]);
  const [generating, setGenerating] = useState(false);
  const [error, setError] = useState('');

  const handleGenerate = async () => {
    setGenerating(true);
    setError('');
    setSummary('');
    setKeywords([]);

    try {
      const result = await api.generateSummary(documentId, length);
      
      if (result.error) {
        setError(result.error);
      } else {
        setSummary(result.content);
        setKeywords(result.keywords || []);
      }
    } catch (err) {
      setError('要約生成に失敗しました');
    } finally {
      setGenerating(false);
    }
  };

  return (
    <div style={{ 
      border: '1px solid #ddd', 
      borderRadius: '8px', 
      padding: '20px',
      marginBottom: '20px'
    }}>
      <h3>要約生成</h3>
      <p><strong>ファイル:</strong> {fileName}</p>

      <div style={{ marginBottom: '15px' }}>
        <label style={{ marginRight: '10px' }}>要約の長さ:</label>
        <select 
          value={length} 
          onChange={(e) => setLength(e.target.value as 'short' | 'medium' | 'long')}
          style={{ padding: '5px' }}
        >
          <option value="short">短い</option>
          <option value="medium">中程度</option>
          <option value="long">長い</option>
        </select>
      </div>

      <button
        onClick={handleGenerate}
        disabled={generating}
        style={{
          padding: '10px 20px',
          backgroundColor: generating ? '#ccc' : '#007bff',
          color: 'white',
          border: 'none',
          borderRadius: '4px',
          cursor: generating ? 'not-allowed' : 'pointer',
          marginBottom: '20px'
        }}
      >
        {generating ? '生成中...' : '要約を生成'}
      </button>

      {error && (
        <div style={{ 
          color: 'red', 
          padding: '10px', 
          backgroundColor: '#ffe6e6', 
          borderRadius: '4px',
          marginBottom: '15px'
        }}>
          {error}
        </div>
      )}

      {summary && (
        <div>
          <h4>生成された要約:</h4>
          <div style={{ 
            padding: '15px', 
            backgroundColor: '#f8f9fa', 
            borderRadius: '4px',
            marginBottom: '15px',
            lineHeight: '1.6'
          }}>
            {summary}
          </div>

          {keywords.length > 0 && (
            <div>
              <h4>キーワード:</h4>
              <div style={{ display: 'flex', flexWrap: 'wrap', gap: '8px' }}>
                {keywords.map((keyword, index) => (
                  <span
                    key={index}
                    style={{
                      padding: '4px 8px',
                      backgroundColor: '#e9ecef',
                      borderRadius: '12px',
                      fontSize: '14px'
                    }}
                  >
                    {keyword}
                  </span>
                ))}
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
};