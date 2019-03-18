#!/usr/bin/env python
#
# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

import unittest

def load_tests(loader, tests, pattern):
	return loader.discover('.')

if __name__ == '__main__':
	unittest.main()
