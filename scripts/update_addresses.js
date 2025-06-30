const fs = require('fs');
const path = require('path');

const DEPLOY_DIR = path.join(__dirname, '..', 'contracts', 'broadcast', 'Deploy.s.sol');

// Detect latest chain-id folder (e.g., 31337, 43114…) by modified time
let broadcastDir = null;
try {
  const subdirs = fs.readdirSync(DEPLOY_DIR, { withFileTypes: true })
    .filter(d => d.isDirectory())
    .map(d => path.join(DEPLOY_DIR, d.name));

  if (subdirs.length === 0) throw new Error('No broadcast folders found');

  // pick the most recently modified directory
  broadcastDir = subdirs.sort((a, b) => fs.statSync(b).mtimeMs - fs.statSync(a).mtimeMs)[0];
} catch (err) {
  console.warn('[WARN] Could not determine broadcast directory:', err.message);
  process.exit(0);
}

const BROADCAST_FILE = path.join(broadcastDir, 'run-latest.json');
const OUTPUT_FILE = path.join(__dirname, '..', 'frontend', 'src', 'addresses.json');

if (!fs.existsSync(BROADCAST_FILE)) {
  console.warn(`[WARN] Broadcast file not found: ${BROADCAST_FILE}. Skipping addresses.json update.`);
  process.exit(0);
}

const data = JSON.parse(fs.readFileSync(BROADCAST_FILE, 'utf8'));
// Build mapping expected by the React app
const addresses = {
  token: '',
  staking: '',
  raffle: '',
  vault: ''
};

for (const tx of data.transactions || []) {
  switch (tx.contractName) {
    case 'ZicoToken':
      addresses.token = tx.contractAddress;
      break;
    case 'ZicoStakingShares':
    case 'ZicoStaking':
      addresses.staking = tx.contractAddress;
      break;
    case 'TreasuryVault':
      addresses.vault = tx.contractAddress;
      break;
    case 'ZicoRaffle':
      addresses.raffle = tx.contractAddress;
      break;
  }
}

fs.writeFileSync(OUTPUT_FILE, JSON.stringify(addresses, null, 2));
console.log('✅ Contract addresses written to', OUTPUT_FILE);

// Remove the browser-only sample code accidentally pasted below.
// END OF SCRIPT 