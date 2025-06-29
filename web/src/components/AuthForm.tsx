import React, { useState } from 'react';
import { api } from '../api/client';

interface AuthFormProps {
  onSuccess: (userId: string) => void;
}

export const AuthForm: React.FC<AuthFormProps> = ({ onSuccess }) => {
  const [isLogin, setIsLogin] = useState(true);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      const result = isLogin 
        ? await api.login(email, password)
        : await api.register(email, password, name);

      if (result.error) {
        setError(result.error);
      } else {
        onSuccess(result.user_id);
      }
    } catch (err) {
      console.error('Auth error:', err);
      if (err instanceof Error) {
        if (err.message.includes('Failed to fetch')) {
          setError('サーバーに接続できません。しばらくお待ちください。');
        } else if (err.message.includes('HTTP error')) {
          setError('サーバーエラーが発生しました。');
        } else {
          setError('ネットワークエラーが発生しました。');
        }
      } else {
        setError('不明なエラーが発生しました。');
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-indigo-500 to-purple-600 p-4">
      <div className="bg-white p-10 rounded-xl shadow-2xl w-full max-w-md">
        <h2 className="text-4xl font-extrabold text-center text-gray-800 mb-8">QuillDeck</h2>
        <p className="text-center text-gray-600 mb-8">
          {isLogin ? 'アカウントにログイン' : '新しいアカウントを作成'}
        </p>
        
        <form onSubmit={handleSubmit}>
          <div className="mb-6">
            <label className="block text-gray-700 text-sm font-bold mb-2">メールアドレス</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="shadow-sm appearance-none border rounded-lg w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-400 focus:border-transparent"
              placeholder="your@example.com"
            />
          </div>

          <div className="mb-6">
            <label className="block text-gray-700 text-sm font-bold mb-2">パスワード</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              className="shadow-sm appearance-none border rounded-lg w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-400 focus:border-transparent"
              placeholder="••••••••"
            />
          </div>

          {!isLogin && (
            <div className="mb-6">
              <label className="block text-gray-700 text-sm font-bold mb-2">名前</label>
              <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
                className="shadow-sm appearance-none border rounded-lg w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-400 focus:border-transparent"
                placeholder="あなたの名前"
              />
            </div>
          )}

          {error && (
            <div className="mb-6 p-3 bg-red-100 border border-red-400 text-red-700 rounded-lg">
              {error}
            </div>
          )}

          <button
            type="submit"
            disabled={loading}
            className={`w-full py-3 px-4 rounded-lg font-bold focus:outline-none focus:shadow-outline transition-colors duration-200 shadow-md ${
              loading 
                ? 'bg-gray-400 cursor-not-allowed text-white' 
                : 'bg-indigo-600 hover:bg-indigo-700 text-white'
            }`}
          >
            {loading ? '処理中...' : (isLogin ? 'ログイン' : 'アカウント作成')}
          </button>
        </form>

        <p className="text-center text-gray-600 text-sm mt-6">
          {isLogin ? 'アカウントをお持ちでないですか？' : 'すでにアカウントをお持ちですか？'}
          <button
            type="button"
            onClick={() => setIsLogin(!isLogin)}
            className="text-indigo-600 hover:underline font-semibold ml-1"
          >
            {isLogin ? '新規登録' : 'ログイン'}
          </button>
        </p>
      </div>
    </div>
  );
};