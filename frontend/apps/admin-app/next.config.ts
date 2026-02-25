import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  transpilePackages: ['shared'],
  webpack: (config) => {
    config.ignoreWarnings = [
      { module: /node_modules/ },
      /data-np-autofill-form-type/,
      /data-np-watching/,
    ];
    return config;
  },
};

export default nextConfig;
