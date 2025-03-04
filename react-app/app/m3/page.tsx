// App.tsx
"use client";
import React, { useState, ChangeEvent } from 'react';
import { useSearchParams } from 'next/navigation';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

// Register Chart.js components
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

// Define a TypeScript interface for the JSON data structure
interface LatencyData {
  [key: string]: number[];
}

const App: React.FC = () => {
  const searchParams = useSearchParams();
  const file = searchParams.get("file");

  // Load JSON data based on the query parameter
  let jsonData;
  if (file === "localToGee") {
    jsonData = require('../[data]/m3/localToGee.json');
  } else if (file === "localToPi") {
    jsonData = require('../[data]/m3/localToPi.json');
  } else if (file === "GeeToPi") {
    jsonData = require('../[data]/m3/GeeToPi.json');
  } else {
    // Default to localToGee if no file parameter is provided
    jsonData = require('../[data]/m3/localToGee.json');
  }

  const data: LatencyData = jsonData;

  // Extract keys from the JSON data (e.g., "256", "512", etc.)
  const keys: string[] = Object.keys(data);
  const [selectedKey, setSelectedKey] = useState<string>(keys[0]);

  // Get the latency data for the selected key
  const latencies: number[] = data[selectedKey];

  // Prepare the chart data
  const chartData = {
    labels: latencies.map((_, index) => `Transmission ${index + 1}`),
    datasets: [
      {
        label: `Latency for Size: ${selectedKey} in ${file}`,
        data: latencies,
        fill: false,
        borderColor: 'rgba(75,192,192,1)',
        backgroundColor: 'rgba(75,192,192,0.4)',
      },
    ],
  };

  // Chart options for responsiveness and styling
  const chartOptions = {
    responsive: true,
    plugins: {
      legend: { position: 'top' as const },
      title: { display: true, text: `UDP RTT Visualization: ${file}` },
    },
  };

  // Event handler for select dropdown
  const handleKeyChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setSelectedKey(e.target.value);
  };

  return (
    <div style={{ width: '80%', margin: '50px auto' }}>
      <h1>UDP RTT Graph: {file}</h1>
      <div style={{ marginBottom: '20px' }}>
        <label>
          Select Data Size:&nbsp;
          <select value={selectedKey} onChange={handleKeyChange}>
            {keys.map((key) => (
              <option key={key} value={key}>
                {key}
              </option>
            ))}
          </select>
        </label>
      </div>
      <Line data={chartData} options={chartOptions} />
    </div>
  );
};

export default App;