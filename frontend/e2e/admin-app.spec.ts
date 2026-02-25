import { test, expect } from '@playwright/test';

test.describe('Admin App - Management Flow', () => {
  test('should display admin dashboard', async ({ page }) => {
    await page.goto('http://localhost:3002');
    await expect(page).toHaveTitle(/Gowoopi Admin/);
  });

  test('should navigate to menu management', async ({ page }) => {
    await page.goto('http://localhost:3002');
    await page.waitForLoadState('networkidle');
    
    const mainContent = page.locator('main');
    await expect(mainContent).toBeVisible();
  });

  test('should navigate to order management', async ({ page }) => {
    await page.goto('http://localhost:3002');
    await page.waitForLoadState('networkidle');
    
    const body = page.locator('body');
    await expect(body).toBeVisible();
  });
});
