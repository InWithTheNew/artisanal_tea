import './App.css';
import React, { useState, useEffect } from 'react';
import { CommandBarField } from './components/command-bar/commandbar';
import { Dropdown, DropdownOption } from './components/dropdown/dropdown';
import { Navbar } from './components/navbar/navbar';
import { Button } from './components/button/button';
import { submitFormData } from './utils/api';
import { fetchDropdownOptions } from './utils/dropdownApi';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL || '';
const APP_URL = `${BACKEND_URL}${process.env.REACT_APP_LIST_APPS_PATH || ''}`; 
const ENV_URL = `${BACKEND_URL}${process.env.REACT_APP_LIST_ENVS_PATH || ''}`;
const SUBMIT_URL = `${BACKEND_URL}${process.env.REACT_APP_SUBMIT_PATH || ''}`;


function App() {
  const [appOptions, setAppOptions] = useState<DropdownOption[]>([]);
  const [selectedApp, setSelectedApp] = useState('');
  const [command, setCommand] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);
  const [response, setResponse] = useState<string | null>(null);

  useEffect(() => {
    fetchDropdownOptions(APP_URL)
      .then(setAppOptions)
      .catch(() => setAppOptions([]));
  }, []);

  const handleSubmit = async () => {
    setLoading(true);
    setError(null);
    setSuccess(false);
    setResponse(null);
    try {
      const res = await submitFormData({
        Name: selectedApp,
        Command: command,
      });
      const data = await res.json();
      if (data.result) {
        setResponse(data.result);
        setSuccess(true);
      } else if (data.error) {
        setError(data.error);
        setResponse(data.error);
      } else {
        setResponse('Unknown response');
      }
    } catch (e: any) {
      setError(e.message || 'Failed to submit');
      setResponse(e.message || 'Failed to submit');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <Navbar />
        <Dropdown
          placeholder="App"
          options={appOptions}
          value={selectedApp}
          onSelect={setSelectedApp}
        />
        <CommandBarField value={command} onChange={setCommand} />
        <Button onClick={handleSubmit} type="submit" disabled={loading}>
          {loading ? 'Submitting...' : 'Submit'}
        </Button>
        {error && <div style={{ color: 'red' }}>{error}</div>}
        {response && (
          <div style={{
            background: '#222',
            color: '#fff',
            padding: '1rem',
            marginTop: '1rem',
            borderRadius: '6px',
            maxWidth: 600,
            wordBreak: 'break-word',
            whiteSpace: 'pre-wrap', // ensures wrapping
            overflowX: 'auto',
          }}>
            <strong>Response:</strong>
            <div style={{whiteSpace: 'pre-wrap'}}>{response}</div>
          </div>
        )}
        {success && <div style={{ color: 'green' }}>Submitted successfully!</div>}
      </header>
    </div>
  );
}

export default App;
