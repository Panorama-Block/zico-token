#!/usr/bin/env bash

# -----------------------------------------------------------------------------
# Export compiled ABIs from the Foundry contracts into the React frontend.
# -----------------------------------------------------------------------------
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
ABIS_DIR="$ROOT_DIR/frontend/src/abis"
CONTRACTS=(ZicoToken ZicoStaking ZicoRaffle TreasuryVault)

mkdir -p "$ABIS_DIR"

for CONTRACT in "${CONTRACTS[@]}"; do
  echo " ➡️  Exporting ABI for $CONTRACT"
  # Determine actual contract name inside the file (special case for staking)
  TARGET_NAME="$CONTRACT"
  if [[ "$CONTRACT" == "ZicoStaking" ]]; then
    TARGET_NAME="ZicoStakingShares"
  fi

  # Use contracts directory as forge root so remappings/libs resolve
  pushd "$ROOT_DIR/contracts" >/dev/null
  forge inspect --json "src/$CONTRACT.sol:$TARGET_NAME" abi > "$ABIS_DIR/$CONTRACT.json"
  popd >/dev/null
  # Keep the output minimal (array of fragments) that ethers.js can consume
done
echo "✅ ABIs exported to $ABIS_DIR" 