import React from 'react';

const TokenIcon = ({ symbol, size = 32, className = "" }) => {
  const getTokenColor = (symbol) => {
    const colors = {
      'ZICO': 'from-purple-600 to-pink-600',
      'ETH': 'from-blue-500 to-purple-600',
      'AVAX': 'from-red-500 to-red-600',
      'MATIC': 'from-purple-500 to-blue-600',
      'ARB': 'from-blue-400 to-blue-600',
      'LINK': 'from-blue-600 to-blue-800',
    };
    return colors[symbol] || 'from-gray-400 to-gray-600';
  };

  return (
    <div 
      className={`inline-flex items-center justify-center rounded-full bg-gradient-to-r ${getTokenColor(symbol)} text-white font-bold text-lg ${className}`}
      style={{ width: size, height: size }}
    >
      {symbol?.charAt(0) || '?'}
    </div>
  );
};

export default TokenIcon; 