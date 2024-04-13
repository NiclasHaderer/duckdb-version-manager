import unittest

import duck_vm


class TestE2E(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        local_versions = duck_vm.list_local_versions(True)
        for version in local_versions:
            duck_vm.uninstall_version(version)

    def test_install_and_removal_version(self):
        remote_versions = duck_vm.list_remote_versions()

        # Install the latest version (at position 0 is nightly)
        latest_version = remote_versions[1]
        duck_vm.install_version(latest_version)

        # Check that the version is installed
        local_versions = duck_vm.list_local_versions()
        self.assertIn(latest_version, local_versions)

        out = duck_vm.run_version(latest_version, "--version")
        self.assertIn(latest_version, out)

        duck_vm.uninstall_version(latest_version)

        local_versions = duck_vm.list_local_versions()
        self.assertNotIn(latest_version, local_versions)

    def test_run_version(self):
        remote_version = "v0.9.1"
        out = duck_vm.run_version(remote_version, "--version")
        self.assertIn(remote_version, out)

    def test_run_and_delete_without_v_prefix(self):
        duck_vm.run_version("0.9.2", "--version")
        local_versions = duck_vm.list_local_versions()
        self.assertIn("v0.9.2", local_versions)
        duck_vm.uninstall_version("0.9.2")

    def test_set_default_version(self):
        duck_vm.set_version_as_default("v0.9.0")
        out = duck_vm.run_default("--version")
        self.assertIn("v0.9.0", out)
        duck_vm.uninstall_version("v0.9.0")
        try:
            duck_vm.run_default("--version")
            self.fail("Should not be able to run default version")
        except FileNotFoundError:
            pass

    def test_run_nightly(self):
        duck_vm.run_version("nightly", "--version")
        duck_vm.uninstall_version("nightly")

    def test_run_invalid_version(self):
        try:
            duck_vm.run_version("not-a-valid-version", "--version")
            self.fail("Should not be able to run invalid version")
        except ValueError as e:
            self.assertIn("exit status 1", str(e))


if __name__ == "__main__":
    unittest.main()
