import React, { useState } from 'react';
import { Connect } from '../wailsjs/go/main/App';

export default function MySQLConnectionForm() {
  const [formData, setFormData] = useState({
    ip: '',
    port: '',
    username: '',
    password: '',
    dbname: ''
  });
  const [connectionStatus, setConnectionStatus] = useState('');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    Connect(formData.ip, formData.port, formData.username, formData.password, formData.dbname).then((result) => {
      if (result) {
        setConnectionStatus('连接成功！');
      } else {
        setConnectionStatus('连接失败，请检查输入信息。');
      }
    })
  };

  return (
    <div className="mx-auto max-w-sm p-4">
      <h1 className="text-2xl font-bold mb-4">MySQL 连接表单</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="ip" className="block text-sm font-medium text-gray-700 mb-1">
            IP 地址
          </label>
          <input
            type="text"
            name="ip"
            onChange={handleChange}
            value={formData.ip}
            required
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
          />
        </div>
        <div>
          <label htmlFor="port" className="block text-sm font-medium text-gray-700 mb-1">
            端口
          </label>
          <input
            type="text"
            name="port"
            onChange={handleChange}
            value={formData.port}
            required
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
          />
        </div>
        <div>
          <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-1">
            用户名
          </label>
          <input
            type="text"
            name="username"
            onChange={handleChange}
            value={formData.username}
            required
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
          />
        </div>
        <div>
          <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
            密码
          </label>
          <input
            type="password"
            name="password"
            onChange={handleChange}
            value={formData.password}
            required
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
          />
        </div>
        <div>
          <label htmlFor="dbname" className="block text-sm font-medium text-gray-700 mb-1">
            数据库名
          </label>
          <input
            type="text"
            name="dbname"
            onChange={handleChange}
            value={formData.dbname}
            required
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          连接
        </button>
      </form>
      {connectionStatus && (
        <div className={`mb-4 ${connectionStatus === '连接成功！' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`}>
          <p>{connectionStatus}</p>
        </div>
      )}
    </div>
  );
};
