const uuid = require("uuid");
const { 
  KEBConfig,
  KEBClient,
  provisionSKR,
  deprovisionSKR,
} = require("../kyma-environment-broker");
const {
  DirectorConfig,
  DirectorClient,
  addScenarioInCompass,
  assignRuntimeToScenario,
  removeScenarioFromCompass,
} = require("../compass");
const {
  GardenerConfig,
  GardenerClient
} = require("../gardener");
const {
  ensureCommerceMockWithCompassTestFixture,
  checkAppGatewayResponse,
  sendEventAndCheckResponse,
} = require("../test/fixtures/commerce-mock");
const {
  debug,
  genRandom,
  initializeK8sClient,
} = require("../utils");

describe("SKR test", function() {
  const keb = new KEBClient(KEBConfig.fromEnv());
  const gardener = new GardenerClient(GardenerConfig.fromEnv());
  const director = new DirectorClient(DirectorConfig.fromEnv());

  const suffix = genRandom(4);
  const appName = `app-${suffix}`;
  const runtimeName = `kyma-${suffix}`;
  const scenarioName = `test-${suffix}`;
  const runtimeID = uuid.v4();

  debug(`RuntimeID ${runtimeID}`, `Scenario ${scenarioName}`, `Runtime ${runtimeName}`, `Application ${appName}`);

  const testNS = "skr-test";

  this.timeout(60 * 60 * 1000 * 2); // 2h
  this.slow(5000);  

  let skr;
  
  it("Provision SKR", async function() {
    skr = await provisionSKR(keb, gardener, runtimeID, runtimeName);
    initializeK8sClient(skr.shoot.kubeconfig);
  });

  it("Assign SKR to scenario", async function() {
    await addScenarioInCompass(director, scenarioName);
    await assignRuntimeToScenario(director, skr.shoot.compassID, scenarioName);
  });

  it("CommerceMock test fixture should be ready", async function () {
    await ensureCommerceMockWithCompassTestFixture(director, appName, scenarioName,  "mocks", testNS);
  });

  it("function should reach Commerce mock API through app gateway", async function () {
    await checkAppGatewayResponse();
  });
    
  it("order.created.v1 event should trigger the lastorder function", async function () {
    await sendEventAndCheckResponse();
  });

  it("Unregister SKR resources from Compass", async function() {
    // removeScenarioFromCompass(director, scenarioName);
  });
   
  it("Deprovision SKR", async function() {
    // await deprovisionSKR(keb, runtimeID);
  });
});