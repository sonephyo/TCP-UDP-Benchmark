// App.tsx
"use client";
import React, { useState, useEffect, ChangeEvent } from 'react';
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
  [outerKey: string]: {
    [innerKey: string]: number[];
  };
}

const App: React.FC = () => {
  const searchParams = useSearchParams();
  const file = searchParams.get("file");

  // Load JSON data based on the query parameter
  let jsonData;
  if (file === "localToGee") {
    jsonData = require('./[data]/m1/localToGee.json');
  } else if (file === "localToPi") {
    jsonData = require('./[data]/m1/localToPi.json');
  } else if (file === "GeeToPi") {
    jsonData = require('./[data]/m1/GeeToPi.json');
  } else {
    // Default to localToGee if no file parameter is provided
    jsonData = require('./[data]/m1/localToGee.json');
  }

  const data: LatencyData = jsonData;

  const mainKeys: string[] = Object.keys(data);
  const [selectedMainKey, setSelectedMainKey] = useState<string>(mainKeys[0]);

  // Extract the inner keys for the selected outer key
  const innerData = data[selectedMainKey];
  const innerKeys: string[] = Object.keys(innerData);
  const [selectedInnerKey, setSelectedInnerKey] = useState<string>(innerKeys[0]);

  // Update the inner key whenever the outer key changes
  useEffect(() => {
    const newInnerKeys: string[] = Object.keys(data[selectedMainKey]);
    setSelectedInnerKey(newInnerKeys[0]);
  }, [selectedMainKey]);

  // Get the latency data from the JSON based on the selected keys
  const latencies: number[] = data[selectedMainKey][selectedInnerKey];

  // Prepare the chart data
  const chartData = {
    labels: latencies.map((_, index) => `Transmission ${index + 1}`),
    datasets: [
      {
        label: `Latency for "${selectedMainKey}" (Size: ${selectedInnerKey}) in ${file}`,
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
      title: { display: true, text: `TCP Latency Visualization: ${file}` },
    },
  };

  // Event handlers for select dropdowns
  const handleMainKeyChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setSelectedMainKey(e.target.value);
  };

  const handleInnerKeyChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setSelectedInnerKey(e.target.value);
  };

  return (
    <div style={{ width: '80%', margin: '50px auto' }}>
      <h1>TCP Latency Graph: {file}</h1>
      <div style={{ marginBottom: '20px' }}>
        <label>
          Select Outer Key:&nbsp;
          <select value={selectedMainKey} onChange={handleMainKeyChange}>
            {mainKeys.map((key) => (
              <option key={key} value={key}>
                {key.length > 30 ? key.substring(0, 30) + '...' : key}
              </option>
            ))}
          </select>
        </label>
      </div>
      <div style={{ marginBottom: '20px' }}>
        <label>
          Select Data Key:&nbsp;
          <select value={selectedInnerKey} onChange={handleInnerKeyChange}>
            {innerKeys.map((innerKey) => (
              <option key={innerKey} value={innerKey}>
                {innerKey}
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