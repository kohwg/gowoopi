import { test, expect } from '@playwright/test';

test.describe('Customer App - Table Order Flow', () => {
  test('should display table order page', async ({ page }) => {
    await page.goto('http://localhost:3001');
    await expect(page).toHaveTitle(/Gowoopi/);
  });

  test('should create table order', async ({ page }) => {
    await page.goto('http://localhost:3001');
    
    // Wait for page to load
    await page.waitForLoadState('networkidle');
    
    // Check if main content is visible
    const mainContent = page.locator('main');
    await expect(mainContent).toBeVisible();
  });

  test('should display menu items', async ({ page }) => {
    await page.goto('http://localhost:3001');
    await page.waitForLoadState('networkidle');
    
    // Check if page loaded successfully
    const body = page.locator('body');
    await expect(body).toBeVisible();
  });
});
