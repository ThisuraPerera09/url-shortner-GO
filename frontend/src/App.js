import React, { useState } from 'react';
import './App.css';
import Header from './components/Header';
import URLShortener from './components/URLShortener';
import URLList from './components/URLList';
import Stats from './components/Stats';

function App() {
  const [activeTab, setActiveTab] = useState('shorten');
  const [refreshList, setRefreshList] = useState(0);
  const [selectedURL, setSelectedURL] = useState(null);

  const handleURLCreated = () => {
    setRefreshList(prev => prev + 1);
  };

  const handleViewStats = (url) => {
    setSelectedURL(url);
    setActiveTab('stats');
  };

  return (
    <div className="App">
      <Header />
      
      <div className="container">
        <nav className="tabs">
          <button 
            className={activeTab === 'shorten' ? 'tab active' : 'tab'}
            onClick={() => setActiveTab('shorten')}
          >
            🔗 Shorten URL
          </button>
          <button 
            className={activeTab === 'list' ? 'tab active' : 'tab'}
            onClick={() => setActiveTab('list')}
          >
            📋 My URLs
          </button>
          <button 
            className={activeTab === 'stats' ? 'tab active' : 'tab'}
            onClick={() => setActiveTab('stats')}
          >
            📊 Analytics
          </button>
        </nav>

        <div className="content">
          {activeTab === 'shorten' && (
            <URLShortener onURLCreated={handleURLCreated} />
          )}
          {activeTab === 'list' && (
            <URLList 
              refresh={refreshList} 
              onViewStats={handleViewStats}
            />
          )}
          {activeTab === 'stats' && (
            <Stats selectedURL={selectedURL} />
          )}
        </div>
      </div>

      <footer className="footer">
        <p>Built with ❤️ using React & Go</p>
        <p>
          <a href="http://localhost:8080/api/health" target="_blank" rel="noopener noreferrer">
            API Status
          </a>
        </p>
      </footer>
    </div>
  );
}

export default App;
